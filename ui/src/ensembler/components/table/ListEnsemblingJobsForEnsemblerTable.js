import React, { Fragment, useEffect, useState } from "react";
import {
  EuiCallOut,
  EuiBasicTable,
  EuiText,
} from "@elastic/eui";
import { DeploymentStatusHealth } from "../../../components/status_health/DeploymentStatusHealth";
import { JobStatus } from "../../../services/job/JobStatus";
import { useConfig } from "../../../config";
import { useTuringApi } from "../../../hooks/useTuringApi";


export const ListEnsemblingJobsForEnsemblerTable = ({
  projectID,
  ensemblerID,
  setDisablePopup,
  disablePopup
}) => {
  const [results, setResults] = useState({ inactiveItems: [], activeItems:[], totalInactiveCount: 0, totalActiveCount:0 });

  const {
    appConfig: {
      tables: { defaultTextSize },
    },
  } = useConfig();

  const [{ data, isLoaded, error }] = useTuringApi(
    `/projects/${projectID}/jobs?ensembler_id=${ensemblerID}`,
    []
  )

  useEffect(() => {
    if (isLoaded && !error) {
      let inactiveItems = data.results.filter((item) => item.status === 'failed_submission' || item.status==='failed_building' || item.status==='failed' || item.status==='completed');
      let activeItems = data.results.filter((item) => item.status !== 'failed_submission' && item.status!=='failed_building' && item.status!=='failed' && item.status!=='completed');
      setResults({
        inactiveItems: inactiveItems ,
        activeItems: activeItems ,
        totalInactiveCount: inactiveItems.length,
        totalActiveCount : activeItems.length
      });
      if (activeItems.length > 0){
        setDisablePopup(true)
      } else {
        setDisablePopup(false)
      }
    }
  }, [data, isLoaded, error, setDisablePopup]);

  const columns = [
    {
      field: "id",
      name: "Id",
      width: "96px",
      render: (id, item) => (
        <EuiText size={defaultTextSize}>
          {id}
        </EuiText>
      ),
    },
    {
      field: "name",
      name: "Name",
      truncateText: true,
      render: (name) => (
        <span className="eui-textTruncate" title={name}>
          {name}
        </span>
      ),
    },
    {
      field: "status",
      name: "Status",
      width: "120px",
      render: (status) => (
        <DeploymentStatusHealth status={JobStatus.fromValue(status)} />
      ),
    },
  ];

  return error ? (
    <EuiCallOut
      title="Sorry, there was an error"
      color="danger"
      iconType="alert">
      <p>{error.message}</p>
    </EuiCallOut>
  ) : (
    <Fragment>
      {disablePopup ? ( results.totalActiveCount > 0 && (
        <div>
          <p>This Ensembler is being used by {results.totalActiveCount} <b>Active Ensembling Jobs</b></p>
          <EuiBasicTable
            items={results.activeItems}
            loading={!isLoaded}
            columns={columns}
            responsive={true}
            tableLayout="auto"
          />
        </div>
      )) : ( results.totalInactiveCount > 0 && (
        <div>
          <p>Deleting this Ensembler will also delete {results.totalInactiveCount} <b>Failed</b> or <b>Completed</b> Ensembling Jobs that use this Ensembler </p>
          <EuiBasicTable
            items={results.inactiveItems}
            loading={!isLoaded}
            columns={columns}
            responsive={true}
            tableLayout="auto"
          />
        </div>
      ))}
    </Fragment>
  );
};