# 23. Prometheus Metrics Naming

Date: 2023-05-22

## Status

Accepted

## Context

Prometheus metrics are used to monitor and measure various aspects of a system. Properly naming these metrics is crucial for clarity and ease of understanding. The current naming convention may not be consistent or follow best practices.

## Decision

To improve the naming of Prometheus metrics, we will adhere to the following best practices:

1. Use descriptive and meaningful names for metrics that accurately represent the measured quantity or behavior.
2. Follow a consistent naming convention throughout the metrics.
3. Utilize labels to enable grouping of related time series.

Please refer to the [Prometheus naming best practices](http://prometheus.io/docs/practices/naming/) for detailed guidelines on naming conventions.

Additionally, leverage the use of labels to provide additional dimensions to the metrics. Labels allow for more fine-grained querying and filtering of metrics data. Please refer to the [Prometheus instrumentation best practices](http://prometheus.io/docs/practices/instrumentation/#use-labels) for guidance on effectively utilizing labels.

## Consequences

By improving the naming of Prometheus metrics and utilizing labels, the following benefits can be expected:

1. Increased clarity and understanding of the metrics, leading to easier interpretation and analysis.
2. Consistent naming conventions will facilitate collaboration among team members and reduce confusion.
3. The use of labels will enable grouping and categorization of related time series, providing better organization and ease of querying.
4. Fine-grained filtering and analysis of metrics data will be possible by leveraging labels.

However, it is important to note that the introduction of these changes may require the following considerations:

1. Existing monitoring systems and dashboards may need to be updated to reflect the new naming conventions and labels.
2. Documentation and training materials should be updated to educate users about the improved practices.
3. Communication and coordination among team members will be necessary to ensure a smooth transition and understanding of the changes.

To mitigate any risks associated with the change, the team should communicate the reasons behind the improvement, provide clear guidelines and examples, and offer support during the transition period. Regular feedback and monitoring should be conducted to address any issues that arise and make further adjustments if needed.
